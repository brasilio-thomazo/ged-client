<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class StoreGroupRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array|string>
     */
    public function rules(): array
    {
        return [
            'name'             => 'required|string|unique:groups,name',
            'privileges'       => 'required|array',
            'types'            => 'required_with:privileges.document|array',
            'departments'      => 'required_with:privileges.document|array',
            'searches'         => 'required_with:privileges.document|array',
            'custom'           => 'array',
            'custom.documents' => 'in:name,department_id,identify',
            'custom.users'     => 'in:name,department_id,identify'
        ];
    }
}
