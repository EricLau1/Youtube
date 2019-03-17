import { INCREMENT ,DECREMENT } from "./types";

export const incr = () => ({ type: INCREMENT });
export const decr = () => ({ type: DECREMENT });