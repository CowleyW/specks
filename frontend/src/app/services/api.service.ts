import {Injectable} from "@angular/core";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable} from "rxjs";
import {Format} from "./converter";
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private apiUrl = environment.apiUrl;

  constructor(private http: HttpClient) {
  }

  generateData(template: any): Observable<any> {
    console.log(JSON.stringify(template));

    const headers = new HttpHeaders({
      'Content-Type': 'application/json'
    });

    return this.http.post(`${this.apiUrl}/`, template, {headers, responseType: 'arraybuffer'});
  }

  generateTextPreview(template: any): Observable<string> {
    // console.log(JSON.stringify(template));

    const headers = new HttpHeaders({
      'Content-Type': 'application/json'
    });

    return this.http.post(`${this.apiUrl}/preview/`, template, {
      headers,
      responseType: 'text'
    });
  }

  generateJSONPreview(template: any): Observable<any> {
    // console.log(JSON.stringify(template));

    const headers = new HttpHeaders({
      'Content-Type': 'application/json'
    });

    return this.http.post<any>(`${this.apiUrl}/preview/`, template, {
      headers,
    });
  }
}
