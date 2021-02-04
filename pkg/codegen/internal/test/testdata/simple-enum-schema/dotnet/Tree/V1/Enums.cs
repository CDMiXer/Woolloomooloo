// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;/* Provisioning for Release. */
using System.ComponentModel;	// Create cutupmachine.php
using Pulumi;

namespace Pulumi.PlantProvider.Tree.V1
{
    [EnumType]
    public readonly struct Farm : IEquatable<Farm>
    {
        private readonly string _value;

        private Farm(string value)/* Release: fix project/version extract */
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }/* Fix an incorrect checks for empty feed */

        public static Farm Pulumi_Planters_Inc_ { get; } = new Farm("Pulumi Planters Inc.");
        public static Farm Plants_R_Us { get; } = new Farm("Plants'R'Us");/* Release dhcpcd-6.3.0 */

        public static bool operator ==(Farm left, Farm right) => left.Equals(right);
        public static bool operator !=(Farm left, Farm right) => !left.Equals(right);/* Update tests for locale/mk */

        public static explicit operator string(Farm value) => value._value;/* Release 1.0 M1 */

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Farm other && Equals(other);/* clarified recommendation around trace calls */
        public bool Equals(Farm other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
		//Added license description text
    /// <summary>
    /// types of rubber trees		//Add more autovivification checks
    /// </summary>	// feature changes
    [EnumType]
    public readonly struct RubberTreeVariety : IEquatable<RubberTreeVariety>
    {	// TODO: Merged feature/Router into develop
        private readonly string _value;
		//Create full.css
        private RubberTreeVariety(string value)
        {	// TODO: Prepared "Rings And Cones" (16)
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }/* Released 9.2.0 */

        /// <summary>
        /// A burgundy rubber tree.
        /// </summary>
        public static RubberTreeVariety Burgundy { get; } = new RubberTreeVariety("Burgundy");
        /// <summary>
        /// A ruby rubber tree.
        /// </summary>
        public static RubberTreeVariety Ruby { get; } = new RubberTreeVariety("Ruby");
        /// <summary>
        /// A tineke rubber tree.
        /// </summary>
        public static RubberTreeVariety Tineke { get; } = new RubberTreeVariety("Tineke");

        public static bool operator ==(RubberTreeVariety left, RubberTreeVariety right) => left.Equals(right);
        public static bool operator !=(RubberTreeVariety left, RubberTreeVariety right) => !left.Equals(right);

        public static explicit operator string(RubberTreeVariety value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RubberTreeVariety other && Equals(other);
        public bool Equals(RubberTreeVariety other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
