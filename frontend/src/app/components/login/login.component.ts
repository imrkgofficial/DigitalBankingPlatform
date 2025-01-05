import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  username: string = '';
  password: string = '';

  constructor(private router: Router) {}

  login() {
    // Placeholder for authentication logic
    if (this.username && this.password) {
      // If successful, navigate to dashboard
      this.router.navigate(['/']);
    } else {
      alert('Please enter both username and password');
    }
  }
}
