import { create } from "zustand";
import axios from "axios";

interface RequestStore {
  url: string;
  jsonBody: any;
  response: any;
  setUrl: (url: string) => void;
  setJsonBody: (jsonBody: any) => void;
  setResponse: (response: any) => void;
  sendRequest: () => void;
}

const useRequestStore = create<RequestStore>((set, get) => ({
  url: "",
  jsonBody: null,
  response: null,
  setUrl: (url: string) => set({ url }),
  setJsonBody: (jsonBody: any) => set({ jsonBody }),
  setResponse: (response: any) => set({ response }),
  sendRequest: async () => {
    const { url, jsonBody } = get();
    console.log(url, jsonBody);
    const data = await axios.post(url, jsonBody);
    set({ response: data });
  },
}));

export default useRequestStore;
