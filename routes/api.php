<?php

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

use App\Http\Controllers\DashController;
use App\Http\Controllers\MailController;
use App\Http\Controllers\ProjectController;
use App\Http\Controllers\SkillController;
use Illuminate\Support\Facades\Route;

Route::get('bg', [DashController::class, 'getBackground'])
    ->middleware('throttle:20,1');

Route::post('mail', [MailController::class, 'sendmail'])
    ->middleware('throttle:10,1');

Route::get('skills', [SkillController::class, 'index'])
    ->middleware('throttle:20,1');

Route::get('projects', [ProjectController::class, 'index'])
    ->middleware('throttle:20,1');
