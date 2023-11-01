<?php

use Illuminate\Support\Facades\Route;

Route::post('/login', [\App\Http\Controllers\ProfileController::class, 'login'])->name('login');
Route::get('/me', [\App\Http\Controllers\ProfileController::class, 'me'])->middleware('auth:sanctum')->name('me');

Route::apiResource('user', \App\Http\Controllers\UserController::class)->middleware('auth:sanctum');
Route::apiResource('group', \App\Http\Controllers\GroupController::class)->middleware('auth:sanctum');

Route::apiResource('document_type', \App\Http\Controllers\DocumentTypeController::class)->middleware('auth:sanctum');
Route::apiResource('department', \App\Http\Controllers\DepartmentController::class)->middleware('auth:sanctum');
Route::apiResource('search', \App\Http\Controllers\SearchController::class)->middleware('auth:sanctum');

Route::apiResource('document', \App\Http\Controllers\DocumentController::class)->middleware('auth:sanctum');
Route::get('document_search', [\App\Http\Controllers\DocumentController::class, 'search'])->middleware('auth:sanctum')->name('document.search');
Route::get("view/{path}", [\App\Http\Controllers\DocumentController::class, 'view'])->where('path', '^[a-z0-9-]+/[a-z0-9-]{36}/[a-z0-9-]{36}\.pdf')->name('storage.temporary');
