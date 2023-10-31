<?php

use App\Http\Controllers\DocumentImageController;
use App\Http\Controllers\DepartmentController;
use App\Http\Controllers\DocumentController;
use App\Http\Controllers\DocumentTypeController;
use App\Http\Controllers\GroupController;
use App\Http\Controllers\ProfileController;
use App\Http\Controllers\SearchController;
use App\Http\Controllers\UserController;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "web" middleware group. Make something great!
|
*/

Route::get('/', fn () => view('app'))->name('home');
Route::post('/login', [ProfileController::class, 'login'])->name('auth.login');
Route::get('/me', [ProfileController::class, 'me'])->middleware(['auth:web'])->name('profile.me');
Route::post('/logout', [ProfileController::class, 'logout'])->middleware(['auth:web'])->name('auth.logout');

Route::controller(GroupController::class)->group(function () {
    Route::get("/group", 'index')->name('group.web.index');
    Route::post("/group", 'store')->name('group.web.store');
    Route::delete("/group", 'destroy')->name('group.web.destroy');
    Route::get("/group/{group}", 'show')->name('group.web.show');
    Route::put("/group/{group}", 'update')->name('group.web.update');
});

Route::controller(UserController::class)->group(function () {
    Route::get("/user", 'index')->name('user.web.index');
    Route::post("/user", 'store')->name('user.web.store');
    Route::delete("/user", 'destroy')->name('user.web.destroy');
    Route::get("/user/{user}", 'show')->name('user.web.show');
    Route::put("/user/{user}", 'update')->name('user.web.update');
});

Route::controller(SearchController::class)->group(function () {
    Route::get("/search", 'index')->name('search.web.index');
    Route::post("/search", 'store')->name('search.web.store');
    Route::delete("/search", 'destroy')->name('search.web.destroy');
    Route::get("/search/{user}", 'show')->name('search.web.show');
    Route::put("/search/{user}", 'update')->name('search.web.update');
});

Route::controller(DocumentTypeController::class)->group(function () {
    Route::get("/type", 'index')->name('type.web.index');
    Route::post("/type", 'store')->name('type.web.store');
    Route::delete("/type", 'destroy')->name('type.web.destroy');
    Route::get("/type/{user}", 'show')->name('type.web.show');
    Route::put("/type/{user}", 'update')->name('type.web.update');
});


Route::controller(DocumentTypeController::class)->group(function () {
    Route::get("/document-type", 'index')->name('document.type.web.index');
    Route::post("/document-type", 'store')->name('document.type.web.store');
    Route::delete("/document-type", 'destroy')->name('document.type.web.destroy');
    Route::get("/document-type/{user}", 'show')->name('document.type.web.show');
    Route::put("/document-type/{user}", 'update')->name('document.type.web.update');
});

Route::controller(DepartmentController::class)->group(function () {
    Route::get("/department", 'index')->name('department.web.index');
    Route::post("/department", 'store')->name('department.web.store');
    Route::delete("/department", 'destroy')->name('department.web.destroy');
    Route::get("/department/{user}", 'show')->name('department.web.show');
    Route::put("/department/{user}", 'update')->name('department.web.update');
});

Route::controller(DocumentImageController::class)->group(function () {
    Route::get("/image", 'index')->name('image.web.index');
    Route::post("/image", 'store')->name('image.web.store');
    Route::delete("/image", 'destroy')->name('image.web.destroy');
    Route::get("/image/{user}", 'show')->name('image.web.show');
    Route::put("/image/{user}", 'update')->name('image.web.update');
});

Route::post('document/search', [DocumentController::class, 'search']);
Route::get('document/search', [DocumentController::class, 'search'])->middleware(['auth:web']);
Route::get('document/{document}/download', [DocumentController::class, 'download']);
Route::get("view/{path}", [DocumentController::class, 'view'])
    ->where('path', '^[a-z0-9-]+/[a-z0-9-]{36}/[a-z0-9-]{36}\.pdf')
    ->name('local.temporary');
