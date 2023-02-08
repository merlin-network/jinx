// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package state

import (
	"github.com/berachain/stargazer/eth/core/precompile"
	libtypes "github.com/berachain/stargazer/lib/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// `ControllableEventManager` defines a cache EventManager that is controllable (snapshottable
// and registrable). It also supports precompile execution by allowing the caller to native events
// as Eth logs.
type ControllableEventManager interface {
	libtypes.Controllable[string]
	sdk.EventManagerI

	// `BeginPrecompileExecution` begins a precompile execution by setting the logs DB.
	BeginPrecompileExecution(precompile.LogsDB)
	// `EndPrecompileExecution` ends a precompile execution by resetting the logs DB to nil.
	EndPrecompileExecution()
}

// `ControllableMultiStore` defines a cache MultiStore that is controllable (snapshottable and
// registrable). It also supports getting the committed KV store from the MultiStore.
type ControllableMultiStore interface {
	libtypes.Controllable[string]
	storetypes.MultiStore

	// `GetCommittedKVStore` returns the committed KV store from the MultiStore.
	GetCommittedKVStore(storetypes.StoreKey) storetypes.KVStore
}

// `AccountKeeper` defines the expected account keeper.
type AccountKeeper interface {
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	GetModuleAddress(moduleName string) sdk.AccAddress
	GetSequence(sdk.Context, sdk.AccAddress) (uint64, error)
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	HasAccount(ctx sdk.Context, addr sdk.AccAddress) bool
	SetAccount(ctx sdk.Context, account authtypes.AccountI)
	RemoveAccount(ctx sdk.Context, account authtypes.AccountI)
	IterateAccounts(ctx sdk.Context, cb func(account authtypes.AccountI) bool)
}

// `BankKeeper` defines the expected bank keeper.
type BankKeeper interface {
	authtypes.BankKeeper
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string,
		recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress,
		recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
}