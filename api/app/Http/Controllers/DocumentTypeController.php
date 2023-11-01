<?php

namespace App\Http\Controllers;

use App\Http\Requests\StoreDocumentTypeRequest;
use App\Http\Requests\UpdateDocumentTypeRequest;
use App\Models\DocumentType;
use App\Models\User;

class DocumentTypeController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        /**
         * @var User
         */
        $user = auth()->user();
        $types = $user->documentTypes();
        $builder = DocumentType::orderby('name');
        if (count($types)) {
            $builder->whereIn('id', $types);
        }
        return response($builder->get());
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StoreDocumentTypeRequest $request)
    {
        $documentType = new DocumentType($request->all());
        $documentType->save();
        return response($documentType, 201);
    }

    /**
     * Display the specified resource.
     */
    public function show(DocumentType $documentType)
    {
        return response($documentType);
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateDocumentTypeRequest $request, DocumentType $documentType)
    {
        $documentType->update($request->all());
        return response($documentType);
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(DocumentType $documentType)
    {
        $documentType->delete();
        return response('', 204);
    }
}
