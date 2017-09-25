<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class DashController extends Controller
{
    public function __invoke() {
        return view('body');
    }

    public function getBackground() {
        $files = \Storage::disk('local')->files('public/bg');
        return json_encode(str_replace('public', 'storage', $files[mt_rand(0, count($files) - 1)]));
    }
}
