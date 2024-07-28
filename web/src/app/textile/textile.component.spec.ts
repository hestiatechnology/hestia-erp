import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TextileComponent } from './textile.component';

describe('TextileComponent', () => {
  let component: TextileComponent;
  let fixture: ComponentFixture<TextileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TextileComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TextileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
