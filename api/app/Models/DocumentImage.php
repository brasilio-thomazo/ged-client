<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\SoftDeletes;
use Illuminate\Support\Facades\Storage;

class DocumentImage extends Model
{
    use HasFactory, SoftDeletes;

    protected $fillable = [
        'document_id',
        'filename',
        'storage_type',
    ];

    protected $appends = [
        'url',
    ];

    protected $hidden = [];

    protected $casts = [];

    protected $dateFormat = 'U';

    public function getUrlAttribute()
    {
        return Storage::disk($this->disk)->temporaryUrl($this->filename, now()->addMinutes(15));
    }

    public function document(): BelongsTo
    {
        return $this->belongsTo(Document::class);
    }
}
