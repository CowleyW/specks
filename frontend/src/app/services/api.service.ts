import {Injectable} from "@angular/core";
import {HttpClient, HttpHeaders} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  constructor(private http: HttpClient) {
  }

  generateData(template: any) {
    console.log(JSON.stringify(template));

    const headers = new HttpHeaders({
      'Content-Type': 'application/json'
    });

    return this.http.post<any>(`http://localhost:4001/`, template, {headers})
      .subscribe({
        next: (response) => console.log("Success\n", response),
        error: (error) => console.error("Error generating data\n", error)
      });
  }
}
