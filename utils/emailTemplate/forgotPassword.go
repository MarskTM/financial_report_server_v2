package emailTemplate

var ForgotPasswordTemplate = `<!DOCTYPE html>
<html
  xmlns:v="urn:schemas-microsoft-com:vml"
  xmlns:o="urn:schemas-microsoft-com:office:office"
  lang="en"
>
  <head>
    <title></title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <style>
      * {
        box-sizing: border-box;
      }

      body {
        margin: 0;
        padding: 0;
      }

      a[x-apple-data-detectors] {
        color: inherit !important;
        text-decoration: inherit !important;
      }

      #MessageViewBody a {
        color: inherit;
        text-decoration: none;
      }

      p {
        line-height: inherit;
      }

      .desktop_hide,
      .desktop_hide table {
        mso-hide: all;
        display: none;
        max-height: 0px;
        overflow: hidden;
      }

      .image_block img + div {
        display: none;
      }

      @media (max-width: 690px) {
        .desktop_hide table.icons-inner,
        .social_block.desktop_hide .social-table {
          display: inline-block !important;
        }

        .icons-inner {
          text-align: center;
        }

        .icons-inner td {
          margin: 0 auto;
        }

        .mobile_hide {
          display: none;
        }

        .row-content {
          width: 100% !important;
        }

        .stack .column {
          width: 100%;
          display: block;
        }

        .mobile_hide {
          min-height: 0;
          max-height: 0;
          max-width: 0;
          overflow: hidden;
          font-size: 0px;
        }

        .desktop_hide,
        .desktop_hide table {
          display: table !important;
          max-height: none !important;
        }
      }
    </style>
  </head>

  <body
    style="
      background-color: #37474f;
      margin: 0;
      padding: 0;
      -webkit-text-size-adjust: none;
      text-size-adjust: none;
    "
  >
    <table
      class="nl-container"
      width="100%"
      border="0"
      cellpadding="0"
      cellspacing="0"
      role="presentation"
      style="
        mso-table-lspace: 0pt;
        mso-table-rspace: 0pt;
        background-color: #37474f;
      "
    >
      <tbody>
        <tr>
          <td>
            <table
              class="row row-1"
              align="center"
              width="100%"
              border="0"
              cellpadding="0"
              cellspacing="0"
              role="presentation"
              style="mso-table-lspace: 0pt; mso-table-rspace: 0pt"
            >
              <tbody>
                <tr>
                  <td>
                    <table
                      class="row-content stack"
                      align="center"
                      border="0"
                      cellpadding="0"
                      cellspacing="0"
                      role="presentation"
                      style="
                        mso-table-lspace: 0pt;
                        mso-table-rspace: 0pt;
                        background-color: #b1e5db;
                        color: #000000;
                        width: 670px;
                        margin: 0 auto;
                      "
                      width="670"
                    >
                      <tbody>
                        <tr>
                          <td
                            class="column column-1"
                            width="100%"
                            style="
                              mso-table-lspace: 0pt;
                              mso-table-rspace: 0pt;
                              font-weight: 400;
                              text-align: left;
                              padding-bottom: 5px;
                              padding-top: 5px;
                              vertical-align: top;
                              border-top: 0px;
                              border-right: 0px;
                              border-bottom: 0px;
                              border-left: 0px;
                            "
                          >
                            <table
                              class="image_block block-1"
                              width="100%"
                              border="0"
                              cellpadding="0"
                              cellspacing="0"
                              role="presentation"
                              style="
                                mso-table-lspace: 0pt;
                                mso-table-rspace: 0pt;
                              "
                            ></table>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </td>
                </tr>
              </tbody>
            </table>
            <table
              class="row row-2"
              align="center"
              width="100%"
              border="0"
              cellpadding="0"
              cellspacing="0"
              role="presentation"
              style="mso-table-lspace: 0pt; mso-table-rspace: 0pt"
            >
              <tbody>
                <tr>
                  <td>
                    <table
                      class="row-content stack"
                      align="center"
                      border="0"
                      cellpadding="0"
                      cellspacing="0"
                      role="presentation"
                      style="
                        mso-table-lspace: 0pt;
                        mso-table-rspace: 0pt;
                        background-color: #b1e5db;
                        color: #000000;
                        width: 670px;
                        margin: 0 auto;
                      "
                      width="670"
                    >
                      <tbody>
                        <tr>
                          <td
                            class="column column-1"
                            width="100%"
                            style="
                              mso-table-lspace: 0pt;
                              mso-table-rspace: 0pt;
                              font-weight: 400;
                              text-align: left;
                              padding-bottom: 5px;
                              padding-top: 5px;
                              vertical-align: top;
                              border-top: 0px;
                              border-right: 0px;
                              border-bottom: 0px;
                              border-left: 0px;
                            "
                          >
                            <table
                              class="image_block block-1"
                              width="100%"
                              border="0"
                              cellpadding="0"
                              cellspacing="0"
                              role="presentation"
                              style="
                                mso-table-lspace: 0pt;
                                mso-table-rspace: 0pt;
                              "
                            >
                              <tr>
                                <td class="pad" style="width: 100%">
                                  <div
                                    class="alignment"
                                    align="center"
                                    style="line-height: 10px"
                                  >
                                    <div style="max-width: 670px">
                                      <a
                                        href="#"
                                        style="outline: none"
                                        tabindex="-1"
                                        ><img
                                          src="https://d1oco4z2z1fhwp.cloudfront.net/templates/default/4056/3275432.png"
                                          style="
                                            display: block;
                                            height: auto;
                                            border: 0;
                                            width: 100%;
                                          "
                                          width="670"
                                          alt="reset password"
                                          title="reset password"
                                      /></a>
                                    </div>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <div
                              class="spacer_block block-2"
                              style="
                                height: 40px;
                                line-height: 40px;
                                font-size: 1px;
                              "
                            >
                              &#8202;
                            </div>
                            <table
                              class="paragraph_block block-3"
                              width="100%"
                              border="0"
                              cellpadding="10"
                              cellspacing="0"
                              role="presentation"
                              style="
                                mso-table-lspace: 0pt;
                                mso-table-rspace: 0pt;
                                word-break: break-word;
                              "
                            >
                              <tr>
                                <td class="pad">
                                  <div
                                    style="
                                      width: 80%;
                                      margin: 0 auto;
                                      color: #393d47;
                                      font-family: 'Helvetica Neue', Helvetica,
                                        Arial, sans-serif;
                                      font-size: 16px;
                                      line-height: 120%;
                                      text-align: center;
                                      mso-line-height-alt: 19.2px;
                                    "
                                  >
                                    <p
                                      style="margin: 0; word-break: break-word"
                                    >
                                      <span
                                        >Email này nhằm mục đích xác minh tài
                                        khoản. Có phải bạn đã gửi yêu cầu quên
                                        mật khẩu tài khoản hệ thống phân công
                                        thực tập.
                                      </span>
                                    </p>
                                    <p
                                      style="
                                        margin: 10px;
                                        word-break: break-word;
                                        text-align: center;
                                         font-size: 16px;
                                         font-weight: bold;
                                      "
                                    >
                                      <span>
                                        Vui lòng nhập mã xác thực bên dưới để tiến hành
                                        đổi mật khẩu.</span
                                      >
                                    </p>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <table
                              class="button_block block-4"
                              width="100%"
                              border="0"
                              cellpadding="20"
                              cellspacing="0"
                              role="presentation"
                              style="
                                mso-table-lspace: 0pt;
                                mso-table-rspace: 0pt;
                              "
                            >
                              <tr>
                                <td class="pad">
                                  <div class="alignment" align="center">
                                    <p
                                      style="
                                        text-decoration: none;
                                        display: inline-block;
                                        color: #d6f8f2;
                                        background-color: #37474f;
                                        border-radius: 24px;
                                        width: auto;
                                        border-top: 0px solid transparent;
                                        font-weight: undefined;
                                        border-right: 0px solid transparent;
                                        border-bottom: 0px solid transparent;
                                        border-left: 0px solid transparent;
                                        padding-top: 5px;
                                        padding-bottom: 5px;
                                        font-family: 'Helvetica Neue', Helvetica,
                                          Arial, sans-serif;
                                        font-size: 16px;
                                        text-align: center;
                                        mso-border-alt: none;
                                        word-break: keep-all;
                                      "
                                      ><span
                                        style="
                                          padding-left: 15px;
                                          padding-right: 15px;
                                          font-size: 16px;
                                          display: inline-block;
                                          letter-spacing: 1px;
                                        "
                                        ><span
                                          style="
                                            word-break: break-word;
                                            line-height: 32px;
                                          "
                                          ><strong>{{.FogortCode}}</strong></span
                                        ></span
                                      ></
                                    >
                                  </div>
                                </td>
                              </tr>
                            </table>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </td>
                </tr>
              </tbody>
            </table>
            <table
              class="row row-3"
              align="center"
              width="100%"
              border="0"
              cellpadding="0"
              cellspacing="0"
              role="presentation"
              style="mso-table-lspace: 0pt; mso-table-rspace: 0pt"
            >
              <tbody>
                <tr>
                  <td>
                    <table
                      class="row-content stack"
                      align="center"
                      border="0"
                      cellpadding="0"
                      cellspacing="0"
                      role="presentation"
                      style="
                        mso-table-lspace: 0pt;
                        mso-table-rspace: 0pt;
                        background-color: #1f1f20;
                        color: #000000;
                        width: 670px;
                        margin: 0 auto;
                      "
                      width="670"
                    >
                      <tbody>
                        <tr>
                          <td
                            class="column column-1"
                            width="33.333333333333336%"
                            style="
                              mso-table-lspace: 0pt;
                              mso-table-rspace: 0pt;
                              font-weight: 400;
                              text-align: left;
                              padding-bottom: 5px;
                              padding-top: 5px;
                              vertical-align: top;
                              border-top: 0px;
                              border-right: 0px;
                              border-bottom: 0px;
                              border-left: 0px;
                            "
                          >
                            <div
                              class="spacer_block block-1"
                              style="
                                height: 20px;
                                line-height: 20px;
                                font-size: 1px;
                              "
                            >
                              &#8202;
                            </div>
                            <table
                              class="image_block block-2"
                              width="100%"
                              border="0"
                              cellpadding="25"
                              cellspacing="0"
                              role="presentation"
                              style="
                                mso-table-lspace: 0pt;
                                mso-table-rspace: 0pt;
                              "
                            >
                              <tr>
                                <td class="pad">
                                  <div
                                    class="alignment"
                                    align="left"
                                    style="line-height: 10px"
                                  >
                                    <div style="max-width: 89.33333333333331px">
                                      <a
                                        href="www.example.com"
                                        target="_blank"
                                        style="outline: none"
                                        tabindex="-1"
                                        ><img
                                          src="https://tse3.mm.bing.net/th?id=OIP.UA2ZfTIZIA1_6MvLYtTGwQHaHa&pid=Api&P=0&h=180"
                                          style="
                                            display: block;
                                            height: auto;
                                            border: 0;
                                            width: 100%;
                                          "
                                          width="89.33333333333331"
                                          alt="company logo"
                                          title="company logo"
                                      /></a>
                                    </div>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <div
                              class="spacer_block block-3"
                              style="
                                height: 20px;
                                line-height: 20px;
                                font-size: 1px;
                              "
                            >
                              &#8202;
                            </div>
                          </td>
                          <td
                            class="column column-2"
                            width="33.333333333333336%"
                            style="
                              mso-table-lspace: 0pt;
                              mso-table-rspace: 0pt;
                              font-weight: 400;
                              text-align: left;
                              padding-bottom: 5px;
                              padding-top: 5px;
                              vertical-align: top;
                              border-top: 0px;
                              border-right: 0px;
                              border-bottom: 0px;
                              border-left: 0px;
                            "
                          >
                            <div
                              class="spacer_block block-1"
                              style="
                                height: 20px;
                                line-height: 20px;
                                font-size: 1px;
                              "
                            >
                              &#8202;
                            </div>
                            <table
                              class="heading_block block-2"
                              width="100%"
                              border="0"
                              cellpadding="0"
                              cellspacing="0"
                              role="presentation"
                              style="
                                mso-table-lspace: 0pt;
                                mso-table-rspace: 0pt;
                              "
                            >
                              <tr>
                                <td
                                  class="pad"
                                  style="
                                    padding-left: 20px;
                                    text-align: center;
                                    width: 100%;
                                  "
                                >
                                  <h3
                                    style="
                                      margin: 0;
                                      color: #ffffff;
                                      direction: ltr;
                                      font-family: Helvetica Neue, Helvetica,
                                        Arial, sans-serif;
                                      font-size: 16px;
                                      font-weight: normal;
                                      line-height: 200%;
                                      text-align: left;
                                      margin-top: 0;
                                      margin-bottom: 0;
                                      mso-line-height-alt: 32px;
                                    "
                                  >
                                    <strong>Đội ngũ phát triển</strong>
                                  </h3>
                                </td>
                              </tr>
                            </table>
                            <table
                              class="divider_block block-3"
                              width="100%"
                              border="0"
                              cellpadding="10"
                              cellspacing="0"
                              role="presentation"
                              style="
                                mso-table-lspace: 0pt;
                                mso-table-rspace: 0pt;
                              "
                            >
                              <tr>
                                <td class="pad">
                                  <div class="alignment" align="left">
                                    <table
                                      border="0"
                                      cellpadding="0"
                                      cellspacing="0"
                                      role="presentation"
                                      width="80%"
                                      style="
                                        mso-table-lspace: 0pt;
                                        mso-table-rspace: 0pt;
                                      "
                                    >
                                      <tr>
                                        <td
                                          class="divider_inner"
                                          style="
                                            font-size: 1px;
                                            line-height: 1px;
                                            border-top: 2px solid #bbbbbb;
                                          "
                                        >
                                          <span>&#8202;</span>
                                        </td>
                                      </tr>
                                    </table>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <table
                              class="paragraph_block block-4"
                              width="100%"
                              border="0"
                              cellpadding="0"
                              cellspacing="0"
                              role="presentation"
                              style="
                                mso-table-lspace: 0pt;
                                mso-table-rspace: 0pt;
                                word-break: break-word;
                              "
                            >
                              <tr>
                                <td
                                  class="pad"
                                  style="
                                    padding-bottom: 10px;
                                    padding-left: 20px;
                                    padding-right: 20px;
                                    padding-top: 10px;
                                  "
                                >
                                  <div
                                    style="
                                      color: #ffffff;
                                      font-family: Helvetica Neue, Helvetica,
                                        Arial, sans-serif;
                                      font-size: 12px;
                                      line-height: 150%;
                                      text-align: left;
                                      mso-line-height-alt: 18px;
                                    "
                                  >
                                    <p
                                      style="margin: 0; word-break: break-word"
                                    >
                                      <span
                                        >Dây là một sản phẩm đang trong giai
                                        đoạn phát triển ý tưởng và chưa được
                                        hoàn thiện. Nếu có bất kỳ vấn đề nào bạn
                                        gặp phải, xin vui lòng phản ảnh lại về
                                        email của mình. Xin chân thành cảm
                                        ơn.<br
                                      /></span>
                                    </p>
                                  </div>
                                </td>
                              </tr>
                            </table>
                          </td>
                          <td
                            class="column column-3"
                            width="33.333333333333336%"
                            style="
                              mso-table-lspace: 0pt;
                              mso-table-rspace: 0pt;
                              font-weight: 400;
                              text-align: left;
                              padding-bottom: 5px;
                              padding-top: 5px;
                              vertical-align: top;
                              border-top: 0px;
                              border-right: 0px;
                              border-bottom: 0px;
                              border-left: 0px;
                            "
                          >
                            <div
                              class="spacer_block block-1"
                              style="
                                height: 20px;
                                line-height: 20px;
                                font-size: 1px;
                              "
                            >
                              &#8202;
                            </div>
                            <table
                              class="heading_block block-2"
                              width="100%"
                              border="0"
                              cellpadding="0"
                              cellspacing="0"
                              role="presentation"
                              style="
                                mso-table-lspace: 0pt;
                                mso-table-rspace: 0pt;
                              "
                            >
                              <tr>
                                <td
                                  class="pad"
                                  style="
                                    padding-left: 20px;
                                    text-align: center;
                                    width: 100%;
                                  "
                                >
                                  <h3
                                    style="
                                      margin: 0;
                                      color: #ffffff;
                                      direction: ltr;
                                      font-family: Helvetica Neue, Helvetica,
                                        Arial, sans-serif;
                                      font-size: 16px;
                                      font-weight: normal;
                                      line-height: 200%;
                                      text-align: left;
                                      margin-top: 0;
                                      margin-bottom: 0;
                                      mso-line-height-alt: 32px;
                                    "
                                  >
                                    <strong>Thông tin liên hệ</strong>
                                  </h3>
                                </td>
                              </tr>
                            </table>
                            <table
                              class="divider_block block-3"
                              width="100%"
                              border="0"
                              cellpadding="10"
                              cellspacing="0"
                              role="presentation"
                              style="
                                mso-table-lspace: 0pt;
                                mso-table-rspace: 0pt;
                              "
                            >
                              <tr>
                                <td class="pad">
                                  <div class="alignment" align="left">
                                    <table
                                      border="0"
                                      cellpadding="0"
                                      cellspacing="0"
                                      role="presentation"
                                      width="80%"
                                      style="
                                        mso-table-lspace: 0pt;
                                        mso-table-rspace: 0pt;
                                      "
                                    >
                                      <tr>
                                        <td
                                          class="divider_inner"
                                          style="
                                            font-size: 1px;
                                            line-height: 1px;
                                            border-top: 2px solid #bbbbbb;
                                          "
                                        >
                                          <span>&#8202;</span>
                                        </td>
                                      </tr>
                                    </table>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <table
                              class="paragraph_block block-4"
                              width="100%"
                              border="0"
                              cellpadding="0"
                              cellspacing="0"
                              role="presentation"
                              style="
                                mso-table-lspace: 0pt;
                                mso-table-rspace: 0pt;
                                word-break: break-word;
                              "
                            >
                              <tr>
                                <td
                                  class="pad"
                                  style="
                                    padding-bottom: 10px;
                                    padding-left: 20px;
                                    padding-right: 20px;
                                    padding-top: 10px;
                                  "
                                >
                                  <div
                                    style="
                                      color: #a9a9a9;
                                      font-family: Helvetica Neue, Helvetica,
                                        Arial, sans-serif;
                                      font-size: 14px;
                                      line-height: 120%;
                                      text-align: left;
                                      mso-line-height-alt: 16.8px;
                                    "
                                  >
                                    <p
                                      style="margin: 0; word-break: break-word"
                                    >
                                      <p
                                       
                                        style="
                                          text-decoration: none;
                                          color: #e9e7e7;
                                        "
                                        rel="noopener"
                                        >manhtokim@gmail.com</
                                      >
                                    </p>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <table
                              class="paragraph_block block-5"
                              width="100%"
                              border="0"
                              cellpadding="0"
                              cellspacing="0"
                              role="presentation"
                              style="
                                mso-table-lspace: 0pt;
                                mso-table-rspace: 0pt;
                                word-break: break-word;
                              "
                            >
                            </table>
                            <table
                              class="paragraph_block block-6"
                              width="100%"
                              border="0"
                              cellpadding="0"
                              cellspacing="0"
                              role="presentation"
                              style="
                                mso-table-lspace: 0pt;
                                mso-table-rspace: 0pt;
                                word-break: break-word;
                              "
                            >
                              <tr>
                                <td
                                  class="pad"
                                  style="
                                    padding-bottom: 10px;
                                    padding-left: 20px;
                                    padding-right: 20px;
                                    padding-top: 0px;
                                  "
                                >
                                  <div
                                    style="
                                      color: #a9a9a9;
                                      font-family: Helvetica Neue, Helvetica,
                                        Arial, sans-serif;
                                      font-size: 14px;
                                      line-height: 120%;
                                      text-align: left;
                                      mso-line-height-alt: 16.8px;
                                    "
                                  >
                                    <p
                                      style="margin: 0; word-break: break-word"
                                    >
                                      <a
                                        href="http://www.example.com"
                                        target="_blank"
                                        style="
                                          text-decoration: underline;
                                          color: #e9e7e7;
                                        "
                                        rel="noopener"
                                        >linkedin.com/in/kim-mạnh-164096270</a>
                                    </p>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <table
                              class="social_block block-7"
                              width="100%"
                              border="0"
                              cellpadding="0"
                              cellspacing="0"
                              role="presentation"
                              style="
                                mso-table-lspace: 0pt;
                                mso-table-rspace: 0pt;
                              "
                            >
                              <tr>
                                <td
                                  class="pad"
                                  style="
                                    padding-bottom: 10px;
                                    padding-left: 20px;
                                    padding-right: 10px;
                                    padding-top: 10px;
                                    text-align: left;
                                  "
                                >
                                  <div class="alignment" align="left">
                                    <table
                                      class="social-table"
                                      width="156px"
                                      border="0"
                                      cellpadding="0"
                                      cellspacing="0"
                                      role="presentation"
                                      style="
                                        mso-table-lspace: 0pt;
                                        mso-table-rspace: 0pt;
                                        display: inline-block;
                                      "
                                    >
                                      <tr>
                                        <td style="padding: 0 20px 0 0">
                                          <a
                                            href="https://www.facebook.com"
                                            target="_blank"
                                            ><img
                                              src="https://app-rsrc.getbee.io/public/resources/social-networks-icon-sets/t-circle-default-gray/facebook@2x.png"
                                              width="32"
                                              height="32"
                                              alt="Facebook"
                                              title="facebook"
                                              style="
                                                display: block;
                                                height: auto;
                                                border: 0;
                                              "
                                          /></a>
                                        </td>
                                        <td style="padding: 0 20px 0 0">
                                          <a
                                            href="https://www.twitter.com"
                                            target="_blank"
                                            ><img
                                              src="https://app-rsrc.getbee.io/public/resources/social-networks-icon-sets/t-circle-default-gray/twitter@2x.png"
                                              width="32"
                                              height="32"
                                              alt="Twitter"
                                              title="twitter"
                                              style="
                                                display: block;
                                                height: auto;
                                                border: 0;
                                              "
                                          /></a>
                                        </td>
                                        <td style="padding: 0 20px 0 0">
                                          <a
                                            href="https://www.instagram.com"
                                            target="_blank"
                                            ><img
                                              src="https://app-rsrc.getbee.io/public/resources/social-networks-icon-sets/t-circle-default-gray/instagram@2x.png"
                                              width="32"
                                              height="32"
                                              alt="Instagram"
                                              title="instagram"
                                              style="
                                                display: block;
                                                height: auto;
                                                border: 0;
                                              "
                                          /></a>
                                        </td>
                                      </tr>
                                    </table>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <div
                              class="spacer_block block-8"
                              style="
                                height: 20px;
                                line-height: 20px;
                                font-size: 1px;
                              "
                            >
                              &#8202;
                            </div>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </td>
                </tr>
              </tbody>
            </table>
            <table
              class="row row-4"
              align="center"
              width="100%"
              border="0"
              cellpadding="0"
              cellspacing="0"
              role="presentation"
              style="
                mso-table-lspace: 0pt;
                mso-table-rspace: 0pt;
                background-color: #ffffff;
              "
            ></table>
          </td>
        </tr>
      </tbody>
    </table>
    <!-- End -->
  </body>
</html>
`