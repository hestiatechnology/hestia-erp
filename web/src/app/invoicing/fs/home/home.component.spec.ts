import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FaturaSimplificadaHomeComponent } from './home.component';

describe('HomeComponent', () => {
  let component: FaturaSimplificadaHomeComponent;
  let fixture: ComponentFixture<FaturaSimplificadaHomeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FaturaSimplificadaHomeComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(FaturaSimplificadaHomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
