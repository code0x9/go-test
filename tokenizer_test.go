package main

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestTokenizer(t *testing.T) {
	s := `

<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>NIAM</title>
    <link rel="stylesheet" href="//cdn.jsdelivr.net/npm/bulma@0.8.0/css/bulma.min.css" />
    <script src="//kit.fontawesome.com/d877489320.js" crossorigin="anonymous"></script>
  </head>
  <body>
    <div class="container is-fluid">
      <div class="columns">
        <div class="column is-one-quarter">
          <aside class="menu">
            <p class="menu-label">
              Login
            </p>
            <ul class="menu-list">
              <li><a href="/oauth2/?idp=">Login(IdP Selection)</a></li>
              <li><a href="/oauth2/?idp=nss">NSS</a></li>
              <li><a href="/oauth2/?idp=dummy">dummy</a></li>
            </ul>
            <p class="menu-label">
              Application
            </p>
            <ul class="menu-list">
              <li><a href="/client">Clients</a></li>
            </ul>
            <p class="menu-label">
              API
            </p>
            <ul class="menu-list">
              <li>
                <a href="/spec/openapi.yaml" target="_blank"> OpenAPI <i class="fas fa-external-link-alt"></i></a>
              </li>
              <li>
                <a href="/spec/swaggerui.html" target="_blank">Swagger UI <i class="fas fa-external-link-alt"></i></a>
              </li>
              <li>
                <a href="/spec/rapidoc.html" target="_blank">RapiDoc <i class="fas fa-external-link-alt"></i></a>
              </li>
              <li>
                <a href="/spec/redoc.html" target="_blank">Redoc <i class="fas fa-external-link-alt"></i></a>
              </li>
            </ul>
            <p class="menu-label">
              Discovery
            </p>
            <ul class="menu-list">
              <li>
                <a href="http://localhost:4444/.well-known/jwks.json" target="_blank">JSON Web Keys <i class="fas fa-external-link-alt"></i></a>
              </li>
              <li>
                <a href="http://localhost:4444/.well-known/openid-configuration" target="_blank">OpenID Connect <i class="fas fa-external-link-alt"></i></a>
              </li>
            </ul>
            <p class="menu-label">
              Community
            </p>
            <ul class="menu-list">
              <li>
                <a href="https://GITTTTTTTTT/nucleo/niam" target="_blank"
                  ><i class="fab fa-github"></i><span> GitHub</span></a
                >
              </li>
              <li>
                <a href="https://line-enterprise.slack.com/archives/C9B5LRZB4" target="_blank"
                  ><i class="fab fa-slack-hash"></i><span> Slack #help_nucleo</span></a
                >
              </li>
            </ul>
          </aside>
        </div>
        <div class="column">
          <div class="content">
            
  
  <h4>ClientID: clientId</h4>
  <h4>ClientSecret: clientSecret</h4>
  <pre>{
  &#34;client_id&#34;: &#34;niam-client&#34;,
  &#34;created_at&#34;: &#34;2019-12-31T01:34:23Z&#34;,
  &#34;grant_types&#34;: [
    &#34;authorization_code&#34;,
    &#34;refresh_token&#34;,
    &#34;client_credentials&#34;,
    &#34;implicit&#34;
  ],
  &#34;redirect_uris&#34;: [
    &#34;http://localhost.linecorp.com:8080/oauth2/callback&#34;,
    &#34;http://localhost:5555/callback&#34;,
    &#34;http://localhost:8081/callback&#34;,
    &#34;https://oauthdebugger.com/debug&#34;
  ],
  &#34;response_types&#34;: [
    &#34;token&#34;,
    &#34;code&#34;,
    &#34;id_token&#34;
  ],
  &#34;scope&#34;: &#34;openid offline offline_access profile&#34;,
  &#34;subject_type&#34;: &#34;public&#34;,
  &#34;token_endpoint_auth_method&#34;: &#34;client_secret_basic&#34;,
  &#34;updated_at&#34;: &#34;2019-12-31T01:34:23Z&#34;,
  &#34;userinfo_signed_response_alg&#34;: &#34;none&#34;
}</pre>

          </div>
        </div>
      </div>
    </div>
  </body>
</html>
`

	body := html.NewTokenizer(strings.NewReader(s))

	for {
		token := body.Next()
		switch token {
		case html.ErrorToken:
			return
		case html.TextToken:
			content := strings.TrimSpace(html.UnescapeString(string(body.Text())))
			if len(content) == 0 {
				continue
			}

			for _, prefix := range []string{"ClientID: ", "ClientSecret: "} {
				if strings.HasPrefix(content, prefix) {
					fmt.Printf("%s\n", strings.TrimPrefix(content, prefix))
				}
			}
		}
	}
}
