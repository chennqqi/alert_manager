package template

var EmailTemplate = `
<!--
Style and HTML derived from https://github.com/mailgun/transactional-email-templates
The MIT License (MIT)
Copyright (c) 2014 Mailgun
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
-->
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">

<head>
  <meta name="viewport" content="width=device-width" />
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>{{ .Subject }}</title>
  <style>
    /* -------------------------------------
    GLOBAL
    A very basic CSS reset
------------------------------------- */
    * {
      margin: 0;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      box-sizing: border-box;
      font-size: 14px;
    }

    img {
      max-width: 100%;
    }

    body {
      -webkit-font-smoothing: antialiased;
      -webkit-text-size-adjust: none;
      width: 100% !important;
      height: 100%;
      line-height: 1.6em;
      /* 1.6em * 14px = 22.4px, use px to get airier line-height also in Thunderbird, and Yahoo!, Outlook.com, AOL webmail clients */
      /*line-height: 22px;*/
    }

    /* Let's make sure all tables have defaults */
    table td {
      vertical-align: top;
    }

    /* -------------------------------------
    BODY & CONTAINER
------------------------------------- */
    body {
      background-color: #f6f6f6;
    }

    .body-wrap {
      background-color: #f6f6f6;
      width: 100%;
    }

    .container {
      display: block !important;
      max-width: 600px !important;
      margin: 0 auto !important;
      /* makes it centered */
      clear: both !important;
    }

    .content {
      max-width: 600px;
      margin: 0 auto;
      display: block;
      padding: 20px;
    }

    /* -------------------------------------
    HEADER, FOOTER, MAIN
------------------------------------- */
    .main {
      background-color: #fff;
      border: 1px solid #e9e9e9;
      border-radius: 3px;
    }

    .content-wrap {
      padding: 30px;
    }

    .content-block {
      padding: 0 0 20px;
    }

    .header {
      width: 100%;
      margin-bottom: 20px;
    }

    .footer {
      width: 100%;
      clear: both;
      color: #999;
      padding: 20px;
    }

    .footer p,
    .footer a,
    .footer td {
      color: #999;
      font-size: 12px;
    }

    /* -------------------------------------
    TYPOGRAPHY
------------------------------------- */
    h1,
    h2,
    h3 {
      font-family: "Helvetica Neue", Helvetica, Arial, "Lucida Grande", sans-serif;
      color: #000;
      margin: 40px 0 0;
      line-height: 1.2em;
      font-weight: 400;
    }

    h1 {
      font-size: 32px;
      font-weight: 500;
      /* 1.2em * 32px = 38.4px, use px to get airier line-height also in Thunderbird, and Yahoo!, Outlook.com, AOL webmail clients */
      /*line-height: 38px;*/
    }

    h2 {
      font-size: 24px;
      /* 1.2em * 24px = 28.8px, use px to get airier line-height also in Thunderbird, and Yahoo!, Outlook.com, AOL webmail clients */
      /*line-height: 29px;*/
    }

    h3 {
      font-size: 18px;
      /* 1.2em * 18px = 21.6px, use px to get airier line-height also in Thunderbird, and Yahoo!, Outlook.com, AOL webmail clients */
      /*line-height: 22px;*/
    }

    h4 {
      font-size: 14px;
      font-weight: 600;
    }

    p,
    ul,
    ol {
      margin-bottom: 10px;
      font-weight: normal;
    }

    p li,
    ul li,
    ol li {
      margin-left: 5px;
      list-style-position: inside;
    }

    /* -------------------------------------
    LINKS & BUTTONS
------------------------------------- */
    a {
      color: #348eda;
      text-decoration: underline;
    }

    .btn-primary {
      text-decoration: none;
      color: #FFF;
      background-color: #348eda;
      border: solid #348eda;
      border-width: 10px 20px;
      line-height: 2em;
      /* 2em * 14px = 28px, use px to get airier line-height also in Thunderbird, and Yahoo!, Outlook.com, AOL webmail clients */
      /*line-height: 28px;*/
      font-weight: bold;
      text-align: center;
      cursor: pointer;
      display: inline-block;
      border-radius: 5px;
      text-transform: capitalize;
    }

    /* -------------------------------------
    OTHER STYLES THAT MIGHT BE USEFUL
------------------------------------- */
    .last {
      margin-bottom: 0;
    }

    .first {
      margin-top: 0;
    }

    .aligncenter {
      text-align: center;
    }

    .alignright {
      text-align: right;
    }

    .alignleft {
      text-align: left;
    }

    .clear {
      clear: both;
    }

    /* -------------------------------------
    ALERTS
    Change the class depending on warning email, good email or bad email
------------------------------------- */
    .alert {
      font-size: 16px;
      color: #fff;
      font-weight: 500;
      padding: 20px;
      text-align: center;
      border-radius: 3px 3px 0 0;
    }

    .alert a {
      color: #fff;
      text-decoration: none;
      font-weight: 500;
      font-size: 16px;
    }

    .alert.alert-critical {
      background-color: #E6522C;
    }

    .alert.alert-warning {
      background-color: #D0021B;
    }

    .alert.alert-cleared {
      background-color: #68B90F;
    }

    /* -------------------------------------
    RESPONSIVE AND MOBILE FRIENDLY STYLES
------------------------------------- */
    @media only screen and (max-width: 640px) {
      body {
        padding: 0 !important;
      }

      h1,
      h2,
      h3,
      h4 {
        font-weight: 800 !important;
        margin: 20px 0 5px !important;
      }

      h1 {
        font-size: 22px !important;
      }

      h2 {
        font-size: 18px !important;
      }

      h3 {
        font-size: 16px !important;
      }

      .container {
        padding: 0 !important;
        width: 100% !important;
      }

      .content {
        padding: 0 !important;
      }

      .content-wrap {
        padding: 10px !important;
      }

      .invoice {
        width: 100% !important;
      }
    }
  </style>
</head>

<body itemscope itemtype="http://schema.org/EmailMessage">

  <table class="body-wrap">
    <tr>
      <td></td>
      <td class="container" width="600">
        <div class="content">
          <table class="main" width="100%" cellpadding="0" cellspacing="0">
            <tr>
              {{ if eq .EventType "ACTIVE" }}
              {{ if eq .AlertSeverity "CRITICAL" }}
              <td class="alert alert-critical">
                {{ else if eq .AlertSeverity "WARN" }}
              <td class="alert alert-warning">
                {{ end }}
                {{ else }}
              <td class="alert alert-cleared">
                {{ end }}

                {{ .Header }}

              </td>
            </tr>
            <tr>
              <td class="content-wrap">
                <table width="100%" cellpadding="0" cellspacing="0">
                  <tr>
                    <td class="content-block">
                      <a href='{{ .AlertMgrURL }}' class="btn-primary">View in Alert Manager</a>
                    </td>
                  </tr>
                  <tr>
                    <td class="content-block">
                      {{ range .AlertParams }}
                      <strong>{{ .Name }}: </strong>{{ .Value }}<br />
                      {{ end }}
                    </td>
                  </tr>
                </table>
              </td>
            </tr>
          </table>

          <div class="footer">
            <table width="100%">
              <tr>
                <td class="aligncenter content-block"><a>Sent by Alert Manager</a></td>
              </tr>
              <tr>
                <td class="aligncenter content-block"><a href='{{ .SentAt }}'>Sent at {{ .SentAt }}</a></td>
              </tr>
            </table>
          </div>
        </div>
      </td>
      <td></td>
    </tr>
  </table>

</body>

</html>
`