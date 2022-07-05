<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Casts\Attribute;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Str;

class Project extends Model
{
    public $timestamps = false;
    protected $hidden = ['heading'];
    protected $appends = [
        'pretty_url',
    ];

    function prettyUrl(): Attribute {
        return Attribute::make(
            get: fn () => Str::of($this->url)
                ->replaceMatches('|^(\w+:)?//|', '')
                ->replaceMatches('|^github\.com/|', ''),
        );
    }
}
