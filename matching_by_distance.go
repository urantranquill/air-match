// dating app


func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the redeemInfo
	for _, elem := range genState.RedeemInfoList {
		k.SetRedeemInfo(ctx, elem)
	}

	// Set all the users
	for _, elem := range genState.ExecutionList {
		k.SetExecution(ctx, elem)
	}

var RandFunc = &functions.Overload{
	// operator for 1 param
	Operator: "rand",
	Function: func(args ...ref.Val) ref.Val {
		return types.Double(rand.Float64())
	decls.NewOverload("min_int_double",
		[]*exprpb.Type{decls.Int, decls.Double},
		decls.Double),
	decls.NewOverload("min_double_int",
		[]*exprpb.Type{decls.Double, decls.Int},
		decls.Double),
	},
}
