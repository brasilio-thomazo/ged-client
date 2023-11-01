<?php

namespace App\Http\Controllers;

use App\Http\Requests\StoreDocumentImageRequest;
use App\Http\Requests\UpdateDocumentImageRequest;
use App\Models\DocumentImage;
use App\Models\Document;
use Illuminate\Http\Request;

class DocumentImageController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $builder = DocumentImage::orderBy('id');
        return response($builder->paginate());
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StoreDocumentImageRequest $request)
    {
        $data = $request->validated();
        $document = Document::findOrFail($data['document_id']);
        $path = sprintf('images/%s', $document->id);
        $data['filename'] =  $request->file('file')->store($path, 'public');
        $image = new DocumentImage($data);
        $image->save();
        return response($image, 201);
    }

    public function multiple(Request $request)
    {
        $request->validate([
            'document_id' => 'required|exists:documents,id',
            'files' => 'required|image|mimes:jpeg,jpg,png,svg,pdf,tiff',
        ]);

        $document = Document::findOrFail($request->get('document_id'));
        $path = sprintf('images/%s', $document->id);
        return response($request);
    }

    /**
     * Display the specified resource.
     */
    public function show(DocumentImage $DocumentImage)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateDocumentImageRequest $request, DocumentImage $DocumentImage)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(DocumentImage $DocumentImage)
    {
        //
    }
}
