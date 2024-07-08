import { Component } from '@angular/core';
import { MoviesApiService } from '../movies-api.service';
import { ActivatedRoute, Params } from '@angular/router';
import { filter } from 'rxjs';
import { Media } from '../media';

@Component({
  selector: 'app-viewer',
  standalone: true,
  imports: [],
  templateUrl: './viewer.component.html',
  styleUrl: './viewer.component.css'
})
export class ViewerComponent {
  src: string = "";
  
  movie!: Media; 

  constructor(private movieApiService: MoviesApiService, private route: ActivatedRoute) {

  }

  ngOnInit() {
    this.movie = {Name: "", Link: ""};
    var name = this.route.snapshot.paramMap.get("name");
    var link = this.route.snapshot.paramMap.get("link");

    if (name && link) {
      this.movie.Name = name;
      this.movie.Link = link;

      this.movieApiService.stream_movie(this.movie).then(Promise => {
        this.movieApiService.download_movie(this.movie)
      })
    }


  


  }
}
