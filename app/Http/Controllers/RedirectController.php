<?php

namespace App\Http\Controllers;

use \App\Redirect as Redirect;

class RedirectController extends Controller
{
    public function index($handle): \Illuminate\Routing\Redirector|\Illuminate\Contracts\Foundation\Application|\Illuminate\Http\RedirectResponse {
        $url = Redirect::select('url')
            ->where('handle', '=', $handle)
            ->get()
            ->pluck('url')
            ->pop() ?? url('/');

        return redirect($url);
    }
}
