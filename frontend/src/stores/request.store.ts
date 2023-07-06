import { create } from "zustand";
import axios, { AxiosError } from "axios";
import { isValidHttpUrl, isJsonString } from "@/utils";

export interface RequestStore {
  url: string;
  jsonBody: any;
  response: any;
  setUrl: (url: string) => void;
  setJsonBody: (jsonBody: any) => void;
  setResponse: (response: any) => void;
  sendRequest: () => void;
  isSendingRequest: boolean;
  isValidUrl: boolean;
  isValidJsonBody: boolean;
  sentRequest: boolean;
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
    set({
      isSendingRequest: true,
      sentRequest: true,
      isValidUrl: isValidHttpUrl(url),
      isValidJsonBody: isJsonString(jsonBody),
    });
    if (isValidHttpUrl(url) && isJsonString(jsonBody)) {
      try {
        const data = await axios.post(url, jsonBody);
        set({ response: data });
      } catch (e) {
        const err = e as AxiosError;
        set({ response: err?.response });
      }
    }
    set({ isSendingRequest: false });
  },
  isSendingRequest: false,
  isValidUrl: false,
  isValidJsonBody: false,
  sentRequest: false,
}));

export default useRequestStore;
