<?php

use Illuminate\Http\Request;

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

Route::get('bg', 'DashController@getBackground')
    ->middleware('throttle:20,1');

Route::post('mail', 'MailController@sendmail')
    ->middleware('throttle:2,1');

Route::get('skills', 'SkillController')
    ->middleware('throttle:20,1');

Route::get('projects', 'ProjectController')
    ->middleware('throttle:20,1');
