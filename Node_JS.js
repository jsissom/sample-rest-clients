var http = require("http");

var options = {
  hostname: 'localhost',
  port: 8080,
  path: '/kfs-dev/coa/accounting_periods/close',
  method: 'POST',
  headers: {
      'Content-Type': 'application/json',
      'authorization': 'NSA_this_is_for_you'
  }
};

var body = {
  'description': 'Document man',
  'universityFiscalYear': 2016,
  'universityFiscalPeriodCode': '05'
};

var req = http.request(options, function(res) {
  console.log('Status code ' + res.statusCode);
  res.on('data', function (body) {
    console.log('Body: ' + body);
  });
});

req.on('error', function(e) {
  console.log('problem with request: ' + e.message);
});

req.write(JSON.stringify(body));
req.end();
