<script type="text/javascript" async="" src="./World of Forks_files/MathJax.js.download"></script>
  <script type="text/javascript" src="./World of Forks_files/highlight.min.js.download"></script>
  <script src="./World of Forks_files/go.min.js.download"></script>
  <script src="./World of Forks_files/typescript.min.js.download"></script>
  hljs.configure({
      tabReplace: '  '
    })
    hljs.registerLanguage("graphql", function (e) { return { aliases: ["gql"], k: { keyword: "query mutation subscription|10 type interface union scalar fragment|10 enum on ...", literal: "true false null" }, c: [e.HCM, e.QSM, e.NM, { cN: "type", b: "[^\\w][A-Z][a-z]", e: "\\W", eE: !0 }, { cN: "literal", b: "[^\\w][A-Z][A-Z]", e: "\\W", eE: !0 }, { cN: "variable", b: "\\$", e: "\\W", eE: !0 }, { cN: "keyword", b: "[.]{2}", e: "\\." }, { cN: "meta", b: "@", e: "\\W", eE: !0 }], i: /([;<']|BEGIN)/ } });
    hljs.initHighlightingOnLoad();
  <script type="text/javascript" src="./World of Forks_files/polyfill-7b0f9ccd98.js.download"></script>
  <script type="text/javascript" src="./World of Forks_files/index-e16d25bc2c.js.download"></script>
  <script src="./World of Forks_files/service-worker.js.download"></script>
  
  if ('serviceWorker' in navigator) {
      window.addEventListener('load', function () {
        navigator.serviceWorker.register('/service-worker.js').then(function (reg) {
          reg.onupdatefound = function () {
            var installingWorker = reg.installing;
            installingWorker.onstatechange = function () {
              switch (installingWorker.state) {
                case 'installed':
                  if (navigator.serviceWorker.controller) {
                    console.log('New or updated content is available.');
                  } else {
                    console.log('Content is now available offline!');
                  }
                  break;
                case 'redundant':
                  console.error('The installing service worker became redundant.');
                  break;
              }
            };
          };
        }).catch(function (e) {
          console.error('Error during service worker registration:', e);
        });
      });
    }