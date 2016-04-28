#!/usr/bin/env perl -w
use strict;
use warnings;
use MIME::Base64;

# http://search.cpan.org/~makamaka/JSON/lib/JSON.pm
# Example install using cpanm:
#   sudo cpanm -i JSON
use JSON;

# http://search.cpan.org/~mcrawfor/REST-Client/lib/REST/Client.pm
# Example install using cpanm:
#   sudo cpanm -i REST::Client
use REST::Client;

# Set the request parameters
my %request_body_map = (
  'description' => 'Closing Time',
  'universityFiscalPeriodCode' => '04',
  'universityFiscalYear' => 2010
);

my $host = 'http://financials:8080';

my $client = REST::Client->new(host => $host);

# POST to the incident table
$client->POST('/kfs-dev/coa/accounting_periods/close',
              encode_json(\%request_body_map),
              {'Authorization' => "NSA_this_is_for_you",
               'Content-Type' => 'application/json',
               'Accept' => 'application/json'});

print 'Response: ' . $client->responseContent() . "\n";
print 'Response status: ' . $client->responseCode() . "\n";
foreach ( $client->responseHeaders() ) {
  print 'Header: ' . $_ . '=' . $client->responseHeader($_) . "\n";
}
