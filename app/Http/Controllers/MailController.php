<?php

namespace App\Http\Controllers;

use Illuminate\Support\Facades\Mail;
use Illuminate\Http\Request;

class MailController extends Controller
{
    public function sendmail(Request $request) {
        $request->validate([
            'name' => 'required',
            'email' => 'required|email',
            'text' => 'required'
        ]);

        $data = [
            'name' => $request->name,
            'email' => $request->email,
            'text' => $request->text
        ];

        Mail::send('emails.contact', $data, function($m) use ($data) {
            $m->from(config('contactform.from.email'), config('contactform.from.name'))
                ->to(config('contactform.to.email'), config('contactform.to.name'))
                ->subject(config('contactform.subject'))
                ->replyTo($data['email'], $data['name']);
        });
    }
}
