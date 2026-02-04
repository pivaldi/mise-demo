import { Component, inject } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { CommonModule } from "@angular/common";

type ItemsResponse = {
  dsn_hint?: string;
  items_count?: number;
  items?: Array<{ id: number; name: string; created_at: string }>;
};

@Component({
  selector: "app-root",
  standalone: true,
  imports: [CommonModule],
  template: `
    <h1>mise-demo (Angular)</h1>
    <p>
      Démo Angular <code>ng serve</code> pilotée par <code>mise</code>, qui tape
      sur l'API Go et la DB Postgres.
    </p>

    <div class="card">
      <h2>Endpoints</h2>
      <ul>
        <li>
          <small>Go</small>:
          <a [href]="apiBase + '/health'" target="_blank" rel="noreferrer"
            >{{ apiBase }}/health</a
          >
        </li>
        <li>
          <small>Go</small>:
          <a [href]="apiBase + '/items'" target="_blank" rel="noreferrer"
            >{{ apiBase }}/items</a
          >
        </li>
        <li>
          <small>PHP</small>:
          <a [href]="phpBase" target="_blank" rel="noreferrer">{{ phpBase }}</a>
        </li>
      </ul>
      <button (click)="refresh()">Rafraîchir /items</button>
    </div>

    <div class="card">
      <h2>Résultat /items</h2>
      <ng-container *ngIf="loading; else content">Chargement…</ng-container>
      <ng-template #content>
        <pre>{{ data | json }}</pre>
      </ng-template>
    </div>

    <div class="card">
      <h2>Commandes</h2>
      <p>
        <code>mise run versions</code> • <code>mise run db:init</code> •
        <code>mise run dev</code>
      </p>
    </div>
  `,
})
export class AppComponent {
  private http = inject(HttpClient);

  apiBase = "http://localhost:8090";
  phpBase = "http://localhost:8091/";
  loading = false;
  data: ItemsResponse | { error: string } | null = null;

  constructor() {
    this.refresh();
  }

  refresh() {
    this.loading = true;
    this.http.get<ItemsResponse>(this.apiBase + "/items").subscribe({
      next: (resp) => {
        this.data = resp;
        this.loading = false;
      },
      error: (err) => {
        this.data = { error: String(err?.message ?? err) };
        this.loading = false;
      },
    });
  }
}
