<html>
    <body>
        <div style="background-color: #607d8b; color: #fff; width: 100%; text-align: center;">
            <h3 style="line-height: 48px; display: inline;">{{ config('contactform.subject') }}</h3>
        </div>
        <table style="margin: 1em auto;">
            <tbody style="vertical-align: top">
                <tr>
                    <td>
                        <b>Name:</b>
                    </td>
                    <td>
                        {{ $name }}
                    </td>
                </tr>
                <tr>
                    <td>
                        <b>Email:</b>
                    </td>
                    <td>
                        {{ $email }}
                    </td>
                </tr>
                <tr>
                    <td>
                        <b>Message:</b>
                    </td>
                    <td>
                        {!! nl2br(e($text)) !!}
                    </td>
                </tr>
            </tbody>
        </table>
        <div style="background-color: #607d8b; color: #fff; width: 100%; text-align: center;">
            <a href="https://gabecook.com/">
                <img src="{{ config('app.url') }}/img/logo-70x70.jpg" style="width: 48px;" alt="Gabe Cook Logo">
            </a>
        </div>
    </body>
</html>
