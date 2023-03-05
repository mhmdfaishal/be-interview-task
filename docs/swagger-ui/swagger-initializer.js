window.onload = function () {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    urls: [
      {
        "url": "./swagger-docs/be-interview.json", "name": "BE Interview Task"
      }
    ],
    "urls.primaryName": "BackEnd Interview Task",
    dom_id: '#swagger-ui',
    deepLinking: true,
    validatorUrl: false,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  });

  //</editor-fold>
};
