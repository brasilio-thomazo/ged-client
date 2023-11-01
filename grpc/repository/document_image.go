package repository

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"br.dev.optimus/hermes/dao"
	"br.dev.optimus/hermes/model"
	"br.dev.optimus/hermes/pb"
	"br.dev.optimus/hermes/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *HermesRepositoryDB) DocumentImageStore(ctx context.Context, stream pb.Hermes_DocumentImageStoreServer) error {
	imageDAO := dao.NewDocumentImageDAO(r.Reader, r.Writer)
	documentDAO := dao.NewDocumentDAO(r.Reader, r.Writer)

	in, err := stream.Recv()
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	documentID, err := uuid.Parse(in.GetInfo().DocumentId)
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := documentDAO.FindById(ctx, documentID); err != nil {
		return status.Error(codes.InvalidArgument, "document id not exists")
	}

	disk := in.GetInfo().Disk

	data := &model.DocumentImage{}
	if err := copier.Copy(data, in); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	var uPath string
	uDir := fmt.Sprintf("%s/%s", os.Getenv("APP_PATH"), documentID.String())
	uFile := fmt.Sprintf("%s%s", id.String(), in.GetInfo().Extension)
	uFilename := fmt.Sprintf("%s/%s", uDir, uFile)
	if disk == "s3" {
		uPath = "/tmp"
	} else {
		uPath = fmt.Sprintf("/home/app/public_html/storage/app/%s", uDir)
	}

	file, err := utils.NewFile(uFile, uPath)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if err := file.Create(); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	log.Printf("create file [%s] [%s]", disk, file.Filename)

	image := &model.DocumentImage{
		DocumentID: documentID,
		Disk:       disk,
		Filename:   uFilename,
		Pages:      in.GetInfo().Pages,
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if _, err := writeData(file, in.GetContent()); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	if disk == "s3" {
		if err = uploadAwsS3(file.Filename, uFilename); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	result, err := imageDAO.Store(ctx, image)
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	reply := &pb.DocumentImageReply{}
	if err := copier.Copy(reply, result); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if err := stream.SendAndClose(reply); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func writeData(file *utils.File, data []byte) (int, error) {
	write, err := file.File.Write(data)
	if err != nil {
		return 0, err
	}
	return write, nil
}

func uploadAwsS3(filename, s3key string) error {
	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	svc := s3.New(sess)
	obj := &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(s3key),
		Body:   file,
	}
	if _, err = svc.PutObject(obj); err != nil {
		return err
	}
	return nil
}
