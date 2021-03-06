package lcd

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/irisnet/irishub/client/context"
	"github.com/irisnet/irishub/client/utils"
	"github.com/irisnet/irishub/modules/record"
)

type postRecordReq struct {
	BaseTx      context.BaseTx `json:"base_tx"`   // basic tx info
	Submitter   string         `json:"submitter"` //  Address of the submitter
	Description string         `json:"description"`
	Data        string         `json:"data"` // for onchain
}

func postRecordHandlerFn(cdc *wire.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req postRecordReq
		err := utils.ReadPostBody(w, r, cdc, &req)
		if err != nil {
			return
		}

		cliCtx = utils.InitRequestClictx(cliCtx, r, req.BaseTx.LocalAccountName, req.Submitter)
		txCtx, err := context.NewTxContextFromBaseTx(cliCtx, cdc, req.BaseTx)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		submitter, err := sdk.AccAddressFromBech32(req.Submitter)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		onchainData := req.Data

		var recordHash string
		var dataSize int64
		// --onchain-data has a high priority over --file-path
		if len(onchainData) != 0 {
			dataSize = int64(binary.Size([]byte(onchainData)))
			if dataSize >= record.UploadLimitOfOnchain {
				utils.WriteErrorResponse(w, http.StatusBadRequest,
					fmt.Sprintf("Upload data is too large, max supported data size is %d", record.UploadLimitOfOnchain))
				return
			}

			sum := sha256.Sum256([]byte(onchainData))
			recordHash = hex.EncodeToString(sum[:])
		} else {
			utils.WriteErrorResponse(w, http.StatusBadRequest,
				"--onchain-data is empty and pleae double check this option")
			return
		}

		recordID := record.KeyRecord(recordHash)
		res, err := cliCtx.QueryStore([]byte(recordID), storeName)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		if len(res) != 0 {
			utils.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Record ID '%s' is already existed", recordID))
			return
		}

		submitTime := time.Now().Unix()

		// create the message
		msg := record.NewMsgSubmitRecord(req.Description, submitTime, submitter, recordHash, dataSize, onchainData)
		err = msg.ValidateBasic()
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.SendOrReturnUnsignedTx(w, cliCtx, txCtx, req.BaseTx, []sdk.Msg{msg})
	}
}
