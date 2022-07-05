<?php

namespace App\Http\Controllers;

use \App\Background as Background;

class DashController extends Controller
{
    public function index(): \Illuminate\Contracts\View\Factory|\Illuminate\Contracts\View\View|\Illuminate\Contracts\Foundation\Application {
        return view('body');
    }

    public function getBackground(): \Illuminate\Http\JsonResponse {
        $bg_url = Background::select('url')
            ->get()
            ->pluck('url');
        return response()->json($bg_url);
    }
}
