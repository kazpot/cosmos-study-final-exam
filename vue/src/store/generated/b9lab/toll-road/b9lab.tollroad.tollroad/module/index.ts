// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateRoadOperator } from "./types/tollroad/tx";
import { MsgCreateRoadOperator } from "./types/tollroad/tx";
import { MsgDeleteRoadOperator } from "./types/tollroad/tx";


const types = [
  ["/b9lab.tollroad.tollroad.MsgUpdateRoadOperator", MsgUpdateRoadOperator],
  ["/b9lab.tollroad.tollroad.MsgCreateRoadOperator", MsgCreateRoadOperator],
  ["/b9lab.tollroad.tollroad.MsgDeleteRoadOperator", MsgDeleteRoadOperator],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgUpdateRoadOperator: (data: MsgUpdateRoadOperator): EncodeObject => ({ typeUrl: "/b9lab.tollroad.tollroad.MsgUpdateRoadOperator", value: MsgUpdateRoadOperator.fromPartial( data ) }),
    msgCreateRoadOperator: (data: MsgCreateRoadOperator): EncodeObject => ({ typeUrl: "/b9lab.tollroad.tollroad.MsgCreateRoadOperator", value: MsgCreateRoadOperator.fromPartial( data ) }),
    msgDeleteRoadOperator: (data: MsgDeleteRoadOperator): EncodeObject => ({ typeUrl: "/b9lab.tollroad.tollroad.MsgDeleteRoadOperator", value: MsgDeleteRoadOperator.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
