<?php

namespace App\Http\Controllers;

use \App\Redirect as Redirect;
use Illuminate\Http\Request;

class RedirectController extends Controller
{
    public function __invoke($handle) {
        $url = Redirect::select('url')
            ->where('handle', '=', $handle)
            ->get()
            ->pluck('url')
            ->pop();

        return redirect($url);
    }
}
