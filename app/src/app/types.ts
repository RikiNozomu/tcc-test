export type loginResponse = {
  token: string;
  expiredAt: string;
}

export type userResponse = {
  id: string;
  username: string;
  createdAt: string;
  updatedAt: string;
}