<?php

namespace App\Providers;

use Illuminate\Support\Facades\Storage;
use Illuminate\Support\Facades\URL;
use Illuminate\Support\ServiceProvider;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Register any application services.
     */
    public function register(): void
    {
    }

    /**
     * Bootstrap any application services.
     */
    public function boot(): void
    {
        Storage::buildTemporaryUrlsUsing(function (string $path, $expiration) {
            return URL::temporarySignedRoute('storage.temporary', $expiration, ['path' => $path]);
        });
    }
}
