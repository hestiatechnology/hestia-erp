import { Injectable } from '@angular/core';
import { BehaviorSubject } from "rxjs";

// If a parent with children has icon and path, they will be ignored
export interface Link {
  name: string;
  path: string;
  icon?: string;
  children?: Link[];
}

@Injectable({
  providedIn: 'root'
})
export class NavbarService {
  private titleSubject = new BehaviorSubject<string>('Hestia ERP');
  private linksSubject = new BehaviorSubject<Link[]>([]);

  title$ = this.titleSubject.asObservable();
  links$ = this.linksSubject.asObservable();

  setTitle(title: string) {
    this.titleSubject.next(title);
  }

  setLinks(links: Link[]) {
    this.linksSubject.next(links);
  }
}
