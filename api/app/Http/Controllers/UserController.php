<?php

namespace App\Http\Controllers;

use App\Http\Requests\StoreUserRequest;
use App\Http\Requests\UpdateUserRequest;
use App\Models\User;

class UserController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $user = auth()->user();
        $builder = User::with(['groups', 'department'])->whereNot('username', '=', 'system');
        if ($user->username != "system") $builder->whereNot('username', '=', 'admin');
        return response($builder->get());
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StoreUserRequest $request)
    {
        $fields = [
            'name',
            'department_id',
            'identity',
            'phone',
            'email',
            'username',
            'password'
        ];

        $user = new User($request->only($fields));
        $user->save();
        $user->groups()->attach($request->get('groups', []));
        $user->groups;
        return response($user, 201);
    }

    /**
     * Display the specified resource.
     */
    public function show(User $user)
    {
        return response($user);
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateUserRequest $request, User $user)
    {
        $fields = [
            'name',
            'department_id',
            'identity',
            'phone',
            'email',
            'username',
            'password'
        ];

        $user->update($request->only($fields));
        $user->groups()->sync($request->get('groups', []));
        $user->groups;
        return response($user);
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(User $user)
    {
        $user->delete();
        return response('', 204);
    }
}
