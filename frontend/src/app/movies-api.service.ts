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
    return axios.get("/api/movies/get?name=" + media.Name + "&link=" + encodeURI(media.Link), {headers: {"Accept": "application/json"}})
  }

  stream_movie(media: Media) {
    return axios.get("/api/movies/serve/" + media.Name, {headers: {"Accept": "application/json"}})
  }

  get_sources(media: Media) {
    return axios.get("/api/movies/sources?link=" + encodeURI(media.Link) + "&name=" + media.Name, {headers: {"Accept": "application/json"}})
  }

  constructor() { }
}
