import { Component } from '@angular/core';
import { MoviesApiService } from '../movies-api.service';
import { Axios } from 'axios';
import { Media } from '../media';
import { ActivatedRoute, Router } from '@angular/router';



@Component({
  selector: 'app-home',
  standalone: true,
  imports: [],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {

  movies!: Media[];

  constructor(private movieApiService: MoviesApiService, private router: Router) {
    this.getMovieList()
  }

  getMovieList() {
    this.movieApiService.get_movie_list()
      .then(Response => {
        this.movies = Response.data;
      })
      .catch(error => {
        console.log(error)
      })
  }

  getMovie(movie: Media) {
    this.movieApiService.get_sources(movie).then(Promise => {
      let sources: Media[] = Promise.data;
      console.log(sources);
      this.router.navigate(['viewer', {name: sources[0].Name, link: sources[0].Link} ])
    });

  }

  ngOnInit() {


  }
}
