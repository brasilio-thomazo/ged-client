<?php

namespace App\Http\Controllers;

use App\Http\Requests\StoreDocumentRequest;
use App\Http\Requests\UpdateDocumentRequest;
use App\Models\Document;
use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Storage;
use Imagick;

class DocumentController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $builder = Document::with(['images'])
            ->orderBy('date_document');
        return response($builder->paginate());
    }

    public function search(Request $request)
    {
        /**
         * @var User
         */
        $user = auth()->user();
        $restrictions = $user->restrictions();

        $builder = Document::with(['documentImages', 'documentType', 'department']);

        if (count($restrictions['departments'])) {
            $builder->whereIn('department_id', $restrictions['departments']);
        }

        if (count($restrictions['types'])) {
            $builder->whereIn('document_type_id', $restrictions['types']);
        }

        foreach ($request->all() as $key => $value) {
            if ($value && preg_match('/^(.+_id|identify|register|storage)$/', $key)) {
                $builder->where($key, $value);
                continue;
            }
            if ($value && preg_match('/^name|comment$/', $key)) {
                $builder->where("$key", "iLIKE",  "%{$value}%");
                continue;
            }
        }

        if ($request->get('start_date')) {
            $start = $request->get('start_date');
            $end = $request->get('end_date');
            $builder->whereBetween('date_document', [$start, $end ? $end : now()]);
        }

        $builder->orderBy('date_document', 'desc');

        return response($builder->paginate()->withQueryString());
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StoreDocumentRequest $request)
    {
        $document = new Document($request->all());
        $document->save();
        return response($document, 201);
    }

    /**
     * Display the specified resource.
     */
    public function show(Document $document)
    {
        $document->images;
        return response($document, 200);
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateDocumentRequest $request, Document $document)
    {
        $document->update($request->all());
        return response($document);
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Document $document)
    {
        $document->images()->delete();
        $document->delete();
        return response([], 204);
    }

    public function view(Request $request, string $path)
    {
        if (!$request->hasValidSignature()) abort(401);
        $content = Storage::disk('local')->get($path);
        return response($content, 200, [
            'Content-Type' => 'application/pdf',
            'Content-Disposition' => 'inline; filename="' . basename($path) . '"'
        ]);
    }
}
