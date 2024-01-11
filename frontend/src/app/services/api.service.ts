import {Injectable} from "@angular/core";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable} from "rxjs";
import {Format} from "./converter";

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  constructor(private http: HttpClient) {
  }

  generateData(template: any): Observable<any> {
    console.log(JSON.stringify(template));

    const headers = new HttpHeaders({
      'Content-Type': 'application/json'
    });

    return this.http.post<any>(`http://localhost:4001/`, template, {headers});
  }

  generateTextPreview(template: any): Observable<string> {
    // console.log(JSON.stringify(template));

    const headers = new HttpHeaders({
      'Content-Type': 'application/json'
    });

    return this.http.post(`http://localhost:4001/preview/`, template, {
      headers,
      responseType: 'text'
    });
  }

  generateJSONPreview(template: any): Observable<any> {
    // console.log(JSON.stringify(template));

    const headers = new HttpHeaders({
      'Content-Type': 'application/json'
    });

    return this.http.post<any>(`http://localhost:4001/preview/`, template, {
      headers,
    });
  }
}
