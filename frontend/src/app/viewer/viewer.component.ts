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


  constructor(private movieApiService: MoviesApiService, private route: ActivatedRoute) {

  }

  ngOnInit() {
    var movie: Media = {Name: "", Link: ""};
    var name = this.route.snapshot.queryParamMap.get("name");
    var link = this.route.snapshot.queryParamMap.get("link");

    if (name && link) {
      movie.Name = name;
      movie.Link = link;

      this.movieApiService.stream_movie(movie).then(Promise => {
        if (Promise.statusText == "downloading") {
          console.log("downloading");
        }
      })
    }


  


  }
}
