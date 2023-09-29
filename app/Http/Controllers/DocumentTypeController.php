<?php

namespace App\Http\Controllers;

use App\Http\Requests\StoreDocumentTypeRequest;
use App\Http\Requests\UpdateDocumentTypeRequest;
use App\Models\DocumentType;
use App\Models\Group;
use App\Models\User;

class DocumentTypeController extends Controller
{
    public function __construct()
    {
    }
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
        //
    }

    /**
     * Display the specified resource.
     */
    public function show(DocumentType $DocumentType)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateDocumentTypeRequest $request, DocumentType $DocumentType)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(DocumentType $DocumentType)
    {
        //
    }
}
