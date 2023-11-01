<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::create('document_images', function (Blueprint $table) {
            $table->id();
            $table->foreignUuid('document_id')->constrained();
            $table->string('disk')->default('local');
            $table->string('filename');
            $table->unsignedInteger('pages')->default(1);
            $table->unsignedBigInteger('created_at');
            $table->unsignedBigInteger('updated_at');
            $table->unsignedBigInteger('deleted_at')->nullable();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('document_images');
    }
};
