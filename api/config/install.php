<?php
return [
    'system' => [
        'password' => env('SYSTEM_PASSWORD', 'system')
    ],
    'admin' => [
        'password' => env('ADMIN_PASSWORD', 'admin')
    ],
    'subscribe' => [
        'url' => env('SUBSCRIBE_URL', null)
    ],
];
