import { writable } from 'svelte/store';
import type { destructuredToken } from './types';

export const loggedIn = writable<boolean>(false);

export const token = writable<destructuredToken>(undefined);