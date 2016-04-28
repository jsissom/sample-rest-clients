/*
 * The Kuali Financial System, a comprehensive financial management system for higher education.
 *
 * Copyright 2005-2016 The Kuali Foundation
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package org.kuali.kfs.rest;

import org.apache.http.HttpResponse;
import org.apache.http.client.ClientProtocolException;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.StringEntity;
import org.apache.http.impl.client.DefaultHttpClient;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class AccountingPeriodCloseJob {
    public static void main(String[] args) {
        try {
            DefaultHttpClient httpClient = new DefaultHttpClient();
            HttpPost request = new HttpPost("http://localhost:8080/kfs-dev/coa/accounting_periods/close");
            request.addHeader("accept", "application/json");
            request.addHeader("content-type", "application/json");
            request.addHeader("authorization","NSA_this_is_for_you");

            StringBuilder sb = new StringBuilder();
            sb.append("{");
            sb.append("\"description\":\"Document: The Next Generation\",");
            sb.append("\"universityFiscalYear\": 2016,");
            sb.append("\"universityFiscalPeriodCode\": \"03\"");
            sb.append("}");
            StringEntity data = new StringEntity(sb.toString());
            request.setEntity(data);

            HttpResponse response = httpClient.execute(request);

            System.out.println("Status Code: " + response.getStatusLine().getStatusCode());
            if (response.getStatusLine().getStatusCode() != 200) {
                throw new RuntimeException("Failed : HTTP error code : " + response.getStatusLine().getStatusCode());
            }

            BufferedReader br = new BufferedReader(new InputStreamReader((response.getEntity().getContent())));

            String output;
            System.out.println("Output from Server .... \n");
            while ((output = br.readLine()) != null) {
                System.out.println(output);
            }

            httpClient.getConnectionManager().shutdown();
        } catch (ClientProtocolException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
