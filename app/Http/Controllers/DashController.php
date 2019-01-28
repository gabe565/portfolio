<?php

namespace App\Http\Controllers;

use \App\Background as Background;
use Illuminate\Http\Request;

class DashController extends Controller
{
    public function __invoke() {
        return view('body');
    }

    public function getBackground() {
        $bg_url = Background::select('url')
            ->get()
            ->pluck('url');
        return response()->json($bg_url);
    }
}
