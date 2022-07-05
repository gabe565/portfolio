<?php

namespace App\Http\Controllers;

use App\Models\Skill as Skill;

class SkillController extends Controller
{
    public function index(): \Illuminate\Http\JsonResponse {
        $skills = Skill::select('skills.title', 'skills.rating', 'skill_headings.title as heading')
            ->join('skill_headings', 'skills.heading', '=', 'skill_headings.id')
            ->orderBy('skill_headings.order', 'asc')
            ->orderBy('skills.rating', 'desc')
            ->orderBy('skills.title', 'asc')
            ->get()
            ->groupBy('heading');

        return response()->json($skills);
    }
}
