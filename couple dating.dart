export 'src/generated/pylons/item.pb.dart';
export 'src/generated/pylons/execution.pb.dart';
export 'src/generated/pylons/trade.pb.dart';
abstract class PylonsWallet {
  static PylonsWallet? _instance;

  static PylonsWallet get instance {
    if (_instance == null) {
      throw WalletInitializationNotDone('SDK not initialized');
    }

var MinIntIntFunc = &functions.Overload{
	// operator for 2 param
	Operator: "min_int_int",
	Binary: func(lhs, rhs ref.Val) ref.Val {
		lftInt64 := lhs.Value().(int64)
		rgtInt64 := rhs.Value().(int64)
		if lftInt64 > rgtInt64 {
			return types.Int(rgtInt64)
		}
		return types.Int(lftInt64)
	},
}
