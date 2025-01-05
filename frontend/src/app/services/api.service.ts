import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  constructor(private http: HttpClient) {}

  login(username: string, password: string): Observable<any> {
    return this.http.post('/api/auth/login', { username, password });
  }

  getAccounts(): Observable<any> {
    return this.http.get('/api/accounts');
  }

  getTransactions(): Observable<any> {
    return this.http.get('/api/transactions');
  }
}
