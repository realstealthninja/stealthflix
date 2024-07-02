import { Injectable } from '@angular/core';
import axios from 'axios';
import { Media } from './media';

@Injectable({
  providedIn: 'root'
})
export class MoviesApiService {
  
  get_movie_list() {
    return axios.get("/api/movies/list", {headers: {"Accept": "application/json"}}) 
  }

  get_movie(media: Media) {
    return axios.get("/api/movies/get?name=" + encodeURI(media.Name) + "&link=" + encodeURI(media.Link), {headers: {"Accept": "application/json"}})
  }

  stream_movie(media: Media) {
    return axios.get("/api/movies/get/" + encodeURI(media.Name), {headers: {"Accept": "application/json"}})
  }

  get_sources(media: Media) {
    return axios.get("/api/movies/sources?link=" + encodeURI(media.Link) + "&name=" + encodeURI(media.Name), {headers: {"Accept": "application/json"}})
  }

  download_movie(media: Media) {
    return axios.post("/api/movies/download", media)
  }

  constructor() { }
}
