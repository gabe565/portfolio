<?php

namespace App\Http\Controllers;

use \App\Project as Project;
use Illuminate\Http\Request;

class ProjectController extends Controller
{
    public function __invoke() {
        return response()->json(Project::all());
    }
}
