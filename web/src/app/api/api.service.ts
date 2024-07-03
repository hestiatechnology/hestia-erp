import { Injectable } from '@angular/core';
import { CompanyManagementClient } from "./grpc/Company_managementServiceClientPb";
import { CreateCompanyRequest } from "./grpc/company_management_pb";
import { Metadata } from "grpc-web";
@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor() { }

  getTest() {
    const client = new CompanyManagementClient('http://localhost:8080');
    const req = new CreateCompanyRequest();
    const md: Metadata = { "authorization": "123" }
    client.createCompany(req, md, function (err, response) {
      if (err) {
        console.log(err.code);
        console.log(err.message);
      } else {
        console.log(response.toObject());
      }
      // ...
    });

  }
}
