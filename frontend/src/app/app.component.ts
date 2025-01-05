import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <div>
      <h1>Welcome to Digital Banking</h1>
      <nav>
        <a routerLink="/">Dashboard</a>
        <a routerLink="/login">Login</a>
        <a routerLink="/account">Account</a>
      </nav>
      <router-outlet></router-outlet>
    </div>
  `,
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'digital-banking';
}
