<?php

return [
    'to' => [
        'email' => env('CONTACTFORM_TO_EMAIL', 'example@example.com'),
        'name' => env('CONTACTFORM_TO_NAME', 'Example')
    ],
    'from' => [
        'email' => env('CONTACTFORM_FROM_EMAIL', 'example@example.com'),
        'name' => env('CONTACTFORM_FROM_NAME', 'Contact Form')
    ],
    'subject' => 'New Form Submission',
];
