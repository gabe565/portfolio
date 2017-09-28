<?php

namespace App\Http\Controllers;

use \App\Skill as Skill;
use Illuminate\Http\Request;

class SkillController extends Controller
{
    public function __invoke() {
        $skills = Skill::select('skills.title', 'skills.rating', 'headings.title as heading')
            ->join('headings', 'skills.heading', '=', 'headings.id')
            ->orderBy('headings.order', 'asc')
            ->orderBy('skills.rating', 'desc')
            ->orderBy('skills.title', 'asc')
            ->get()
            ->groupBy('heading');

        $skills = $skills->map(function($item, $key) {
            $count = $item->count();
            if ($count > 1) {
                $chunk = ceil($count / 2);
                return $item->chunk($chunk, false);
            } else {
                return $item;
            }
        });

        return response()->json($skills);
    }
}
